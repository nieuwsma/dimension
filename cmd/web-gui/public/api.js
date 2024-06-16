
async function fetchActiveSessions() {
    try {
        const response = await fetch('http://localhost:8080/training', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        console.log('Active Sessions:', data);  // Debugging line to log the response
        return data.trainingSessions;
    } catch (error) {
        console.error('Error fetching active sessions:', error);
        return [];
    }
}

async function fetchRules() {
    try {
        const response = await fetch('http://localhost:8080/rules', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        console.log('Rules:', data);  // Debugging line to log the response
        return data.Tasks;  // Return the entire task objects
    } catch (error) {
        console.error('Error fetching rules:', error);
        return [];
    }
}


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
        // Convert tasks to an array of objects with Name and Description
        return {
            trainID: data.trainID,
            tasks: data.tasks.map(task => ({ Name: task, Description: '' })) // Add empty description
        };
    } catch (error) {
        console.error('Error fetching tasks:', error);
        return null;
    }
}

async function regenerateTasks(trainID) {
    try {
        const response = await fetch(`http://localhost:8080/training/${trainID}/regenerate`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        console.log('Regenerated Tasks:', data);  // Debugging line to log the response
        // Convert tasks to an array of objects with Name and Description
        return {
            tasks: data.tasks.map(task => ({ Name: task, Description: '' }))
        };
    } catch (error) {
        console.error('Error regenerating tasks:', error);
        return null;
    }
}

