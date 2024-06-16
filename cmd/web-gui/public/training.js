document.addEventListener('DOMContentLoaded', async () => {
    const urlParams = new URLSearchParams(window.location.search);
    const sessionID = urlParams.get('sessionID');

    if (sessionID) {
        document.getElementById('session-id').textContent = `Session ID: ${sessionID}`;
        const tasks = await fetchTasks();  // Fetch tasks for the specific session ID
        console.log('Tasks:', tasks);  // Debugging line to log the tasks
        renderTasks(tasks);
    } else {
        console.error('No session ID found.');
    }
});

async function regenerateTasks() {
    const urlParams = new URLSearchParams(window.location.search);
    const sessionID = urlParams.get('sessionID');

    if (sessionID) {
        const tasks = await regenerateTasks(sessionID);
        console.log('Regenerated Tasks:', tasks);  // Debugging line to log the tasks
        renderTasks(tasks);
    } else {
        console.error('No session ID found for regeneration.');
    }
}
