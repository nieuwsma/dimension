document.addEventListener('DOMContentLoaded', async () => {
    const urlParams = new URLSearchParams(window.location.search);
    const trainID = urlParams.get('trainID');
    console.log('trainID:', trainID); // Debugging line to log trainID

    if (trainID) {
        document.getElementById('train-id').textContent = `Training Session ID: ${trainID}`;
        const response = await fetchTrainingStatistics(trainID);  // Fetch tasks for the specific trainID
        console.log(response)
        console.log("AndRew")

        // Add debugging logs
        console.log('API Response:', response);
        console.log('Response.tasks:', response ? response.tasks : 'No response');
        console.log('Response.turn:', response ? response.turn : 'No response');


        if (response && response.tasks && Array.isArray(response.turn)) {
            const tasks = response.tasks;
            const turns = response.turn;
            const expirationTime = response.expirationTime;
            console.log('Tasks:', tasks);  // Debugging line to log the tasks
            console.log('Turns:', turns);  // Debugging line to log the turns
            console.log('Expiration Time:', expirationTime); // Debugging line to log the expiration time
            renderStatistics(tasks, turns, expirationTime); // Pass tasks, turns, and expirationTime to renderStatistics
        } else {
            console.error('Failed to fetch statistics for the training session.');
        }
    } else {
        console.error('No trainID found.');
    }
});


