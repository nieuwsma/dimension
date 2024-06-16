document.addEventListener('DOMContentLoaded', async () => {
    const urlParams = new URLSearchParams(window.location.search);
    const sessionID = urlParams.get('sessionID');

    if (sessionID) {
        document.getElementById('session-id').textContent = `Session ID: ${sessionID}`;
        const response = await fetchTasks();  // Fetch tasks for the specific session ID
        if (response && response.tasks) {
            const tasks = response.tasks;
            console.log('Tasks:', tasks);  // Debugging line to log the tasks
            renderTasks(tasks);
        } else {
            console.error('Failed to fetch tasks for the session.');
        }
    } else {
        console.error('No session ID found.');
    }
});

async function regenerateTasks() {
    const urlParams = new URLSearchParams(window.location.search);
    const sessionID = urlParams.get('sessionID');

    if (sessionID) {
        const response = await regenerateTasks(sessionID);
        if (response && response.tasks) {
            const tasks = response.tasks;
            console.log('Regenerated Tasks:', tasks);  // Debugging line to log the tasks
            renderTasks(tasks);
        } else {
            console.error('Failed to regenerate tasks for the session.');
        }
    } else {
        console.error('No session ID found for regeneration.');
    }
}
