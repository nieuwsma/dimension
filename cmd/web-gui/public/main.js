document.addEventListener('DOMContentLoaded', async () => {
    const sessions = await fetchActiveSessions();
    console.log('Sessions:', sessions);  // Debugging line to log the sessions

    // Sort sessions alphabetically
    sessions.sort();

    // Render sessions
    renderSessions(sessions);
});

function viewRules() {
    window.location.href = 'rules.html';
}

function startNewSession() {
    fetchTasks().then(data => {
        if (data) {
            console.log('New Session Data:', data);  // Debugging line to log new session data
            const { trainID, tasks } = data;
            console.log('New Session ID:', trainID);  // Debugging line to log new session ID
            console.log('New Session Tasks:', tasks);  // Debugging line to log new session tasks
            // Redirect to training page with new session ID and tasks
            window.location.href = `training.html?sessionID=${trainID}`;
        } else {
            console.error('Failed to create a new session.');
        }
    });
}


function joinSession(sessionID) {
    // Logic to join a session, e.g., redirect to training page with session ID
    window.location.href = `training.html?sessionID=${sessionID}`;
}
