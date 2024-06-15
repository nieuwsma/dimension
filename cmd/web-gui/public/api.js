async function fetchTasks() {
    try {
        const response = await fetch('http://localhost:8080/training', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({})
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        console.log('API Response:', data);  // Debugging line to log the response
        return data.tasks;
    } catch (error) {
        console.error('Error fetching tasks:', error);
        return [];
    }
}
