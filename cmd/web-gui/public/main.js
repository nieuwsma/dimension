document.addEventListener('DOMContentLoaded', async () => {
    const tasks = await fetchTasks();
    console.log('Tasks:', tasks);  // Debugging line to log the tasks
    renderTasks(tasks);
});
