document.addEventListener('DOMContentLoaded', async () => {
    const trainingSessions = await fetchActiveTrainingSessions();
    console.log('Sessions:', trainingSessions);  // Debugging line to log the trainingSessions

    // Sort trainingSessions alphabetically
    trainingSessions.sort();

    // Render trainingSessions
    renderTrainingSessions(trainingSessions);
});

function createGame() {
    createTrainingSession().then(data => {
        if (data) {
            console.log('New Training Session Data:', data);  // Debugging line to log new session data
            const { trainID, tasks } = data;
            console.log('New Training Session ID:', trainID);  // Debugging line to log new session ID
            console.log('New Training Session Tasks:', tasks);  // Debugging line to log new session tasks
            // Redirect to training page with new  trainID and tasks
            window.location.href = `play-game.html?trainID=${trainID}`;
        } else {
            console.error('Failed to create a new game.');
        }
    });
}

function joinGame(trainID) {
    // Logic to join a training session, e.g., redirect to training page with  trainID
    window.location.href = `play-game.html?trainID=${trainID}`;
}
