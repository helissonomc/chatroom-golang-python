const signalingServer = new WebSocket("ws://localhost:8080/ws");

let localStream;
let localPeerConnection;
let remotePeerConnection;

const localVideo = document.getElementById("localVideo");
const remoteVideo = document.getElementById("remoteVideo");

// Get access to the user's video and audio
navigator.mediaDevices.getUserMedia({ video: true, audio: true })
    .then(stream => {
        // Set the local video stream
        localVideo.srcObject = stream;
        localStream = stream;
        // Start the call after the stream is ready
        startCall();
    })
    .catch(error => console.error("Error accessing media devices.", error));

// Handle messages from the signaling server
signalingServer.onmessage = async (message) => {
    const data = JSON.parse(message.data);

    if (data.type === "offer") {
        await handleOffer(data.offer);
    } else if (data.type === "answer") {
        await handleAnswer(data.answer);
    } else if (data.type === "candidate") {
        await handleCandidate(data.candidate);
    }
};

async function startCall() {
    if (!localStream) {
        console.error("Local stream is not available.");
        return;
    }

    // Create peer connections
    localPeerConnection = new RTCPeerConnection();
    remotePeerConnection = new RTCPeerConnection();

    // Handle ICE candidates for local peer
    localPeerConnection.onicecandidate = ({ candidate }) => {
        if (candidate) {
            signalingServer.send(JSON.stringify({ type: "candidate", candidate }));
        }
    };

    // Handle ICE candidates for remote peer
    remotePeerConnection.onicecandidate = ({ candidate }) => {
        if (candidate) {
            signalingServer.send(JSON.stringify({ type: "candidate", candidate }));
        }
    };

    // Display remote stream when received
    remotePeerConnection.ontrack = (event) => {
        remoteVideo.srcObject = event.streams[0];
    };

    // Add local stream tracks to local peer connection
    localStream.getTracks().forEach(track => localPeerConnection.addTrack(track, localStream));

    // Create and send offer from local peer
    const offer = await localPeerConnection.createOffer();
    await localPeerConnection.setLocalDescription(offer);
    signalingServer.send(JSON.stringify({ type: "offer", offer }));
}

async function handleOffer(offer) {
    if (!remotePeerConnection) {
        console.error("Remote peer connection is not available.");
        return;
    }

    // Set remote description and create answer
    await remotePeerConnection.setRemoteDescription(new RTCSessionDescription(offer));
    localStream.getTracks().forEach(track => remotePeerConnection.addTrack(track, localStream));
    const answer = await remotePeerConnection.createAnswer();
    await remotePeerConnection.setLocalDescription(answer);

    // Send answer back to signaling server
    signalingServer.send(JSON.stringify({ type: "answer", answer }));
}

async function handleAnswer(answer) {
    if (!localPeerConnection) {
        console.error("Local peer connection is not available.");
        return;
    }

    // Set remote description on the local peer connection
    await localPeerConnection.setRemoteDescription(new RTCSessionDescription(answer));
}

async function handleCandidate(candidate) {
    try {
        // Add received ICE candidate to both peer connections
        await localPeerConnection.addIceCandidate(candidate);
        await remotePeerConnection.addIceCandidate(candidate);
    } catch (error) {
        console.error("Error adding received ICE candidate", error);
    }
}
