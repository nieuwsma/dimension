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
        try {
            const response = await fetch(`http://localhost:8080/training/${sessionID}/regenerate`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const data = await response.json();
            console.log('Regenerated Tasks:', data);  // Debugging line to log the tasks

            // Convert tasks to an array of objects with Name and Description
            const tasks = data.tasks.map(task => ({ Name: task, Description: '' }));

            // Render the tasks
            renderTasks(tasks);
        } catch (error) {
            console.error('Error regenerating tasks:', error);
        }
    } else {
        console.error('No session ID found for regeneration.');
    }
}
