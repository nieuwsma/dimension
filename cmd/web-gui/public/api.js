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

async function retrieveTrainingSession(ID) {
    const trainID = encodeURIComponent(ID);

    try {
        const response = await fetch(`http://localhost:8080/training/${trainID}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        console.log('API Response:', data);  // Debugging line to log the response
        // Convert tasks to an array of objects with Name and Description
        return {
            trainID: data.trainID,
            tasks: data.tasks.map(task => ({Name: task, Description: ''})) // Add empty description
        };
    } catch (error) {
        console.error('Error fetching tasks:', error);
        return null;
    }
}

async function createTrainingSession() {
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
            tasks: data.tasks.map(task => ({Name: task, Description: ''})) // Add empty description
        };
    } catch (error) {
        console.error('Error fetching tasks:', error);
        return null;
    }
}

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
            const tasks = data.tasks.map(task => ({Name: task, Description: ''}));

            // Render the tasks
            renderTasks(tasks);
        } catch (error) {
            console.error('Error regenerating tasks:', error);
        }
    } else {
        console.error('No session ID found for regeneration.');
    }
}

async function submitTurn(trainID, playerName, slotData) {
    console.log('Submitting turn');  // Debugging line to check function call
    console.log('Train ID:', trainID);  // Debugging line to log trainID
    console.log('Player Name:', playerName);  // Debugging line to log playerName
    console.log('Slot Data:', slotData);  // Debugging line to log slotData

    const payload = slotData;
    const encodedPlayerName = encodeURIComponent(playerName);
    try {
        const response = await fetch(`http://localhost:8080/training/${encodeURIComponent(trainID)}/turn/${encodedPlayerName}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        });

        if (!response.ok) {
            const errorData = await response.json();
            return { error: errorData };
        }

        return await response.json();
    } catch (error) {
        console.error('Error submitting turn:', error);
        return { error: { detail: 'Network error, please try again later.' } };
    }
}


async function fetchRuleDescriptions() {
    try {
        const response = await fetch(`http://localhost:8080/rules`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return Array.isArray(data.Tasks) ? data.Tasks : [];
    } catch (error) {
        console.error('Error fetching rule descriptions:', error);
        return [];
    }
}

