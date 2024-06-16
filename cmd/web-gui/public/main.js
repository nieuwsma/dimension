document.addEventListener('DOMContentLoaded', async () => {
    const sessions = await fetchActiveSessions();
    console.log('Sessions:', sessions);  // Debugging line to log the sessions
    renderSessions(sessions);
});

function viewRules() {
    window.location.href = 'rules.html';
}

function startNewSession() {
    fetchTasks().then(tasks => {
        console.log('New Session Tasks:', tasks);  // Debugging line to log new session tasks
        // Redirect to training page with new session ID and tasks
        window.location.href = 'training.html'; // Update with actual session ID if necessary
    });
}

function joinSession(sessionID) {
    // Logic to join a session, e.g., redirect to training page with session ID
    window.location.href = `training.html?sessionID=${sessionID}`;
}
