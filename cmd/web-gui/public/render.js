function renderTasks(tasks, containerId = 'card-container', includeDescription = false) {
    if (!tasks || tasks.length === 0) {
        console.error('No tasks to render.');
        return;
    }

    const container = document.getElementById(containerId);
    if (!container) {
        console.error(`Container with ID ${containerId} not found.`);
        return;
    }
    container.innerHTML = ''; // Clear previous content

    tasks.forEach(task => {
        console.log('Task:', task);  // Debugging line to log each task

        const {Name, Description} = task;
        const parts = Name.split('-');

        let card = document.createElement('div');
        if (includeDescription) {
            card.className = 'card-with-description';
        } else {
            card.className = 'card';
        }
        // i changed my mind, I dont want to render the hover element
        // const taskName = document.createElement('div');
        // taskName.className = 'task-name';
        // taskName.textContent = Description;
        // card.appendChild(taskName);

        // Create a graphical representation based on the task type
        switch (parts[0]) {
            case 'QUANTITY':
                card.className += ' quantity';
                card.innerHTML += `<div class="circle ${getColorClass(parts[2])}"></div>
                                  <div class="number">${parts[1]}</div>`;
                break;
            case 'TOUCH':
                card.innerHTML += `<div class="side-by-side">
                                     <div class="circle ${getColorClass(parts[1])}"></div>
                                     <div class="circle ${getColorClass(parts[2])}"></div>
                                   </div>`;
                break;
            case 'NOTOUCH':
                card.innerHTML += `<div class="side-by-side notouch">
                                     <div class="circle ${getColorClass(parts[1])}"></div>
                                     <div class="circle ${getColorClass(parts[2])}"></div>
                                     <div class="cross"></div>
                                   </div>`;
                break;
            case 'SUM':
                const color1 = getColorClass(parts[2]);
                const color2 = getColorClass(parts[3]);
                card.innerHTML += `<div class="circle" style="background: linear-gradient(to right, ${color1} 50%, ${color2} 50%);"></div>
                                   <div class="number">${parts[1]}</div>`;
                break;
            case 'GT':
                card.innerHTML += `<div class="side-by-side gt">
                                     <div class="circle ${getColorClass(parts[1])}"></div>
                                     <div class="symbol">></div>
                                     <div class="circle ${getColorClass(parts[2])}"></div>
                                   </div>`;
                break;
            case 'TOP':
                card.className += ' top';
                card.innerHTML += `<div class="circle gradient">
                                      <div class="cross"></div>
                                    </div>
                                  <div class="circle ${getColorClass(parts[1])}"></div>`;
                break;
            case 'BOTTOM':
                card.className += ' bottom';
                card.innerHTML += `<div class="circle ${getColorClass(parts[1])}"></div>
                                  <div class="circle gradient">
                                      <div class="cross"></div>
                                    </div>`;
                break;
            default:
                card.innerHTML = `<p>Unknown Task: ${Name}</p>`;
        }

        if (includeDescription && Description) {
            const descriptionDiv = document.createElement('div');
            descriptionDiv.className = 'description';
            descriptionDiv.textContent = Description;
            card.appendChild(descriptionDiv);
        }

        container.appendChild(card);
    });
}

function getColorClass(colorCode) {
    switch (colorCode) {
        case 'G':
            return 'green';
        case 'B':
            return 'blue';
        case 'K':
            return 'black';
        case 'O':
            return 'orange';
        case 'W':
            return 'white';
        default:
            return 'gray';
    }
}


function renderTrainingSessions(trainingSessions) {
    const container = document.getElementById('session-list');
    container.innerHTML = ''; // Clear previous content

    trainingSessions.forEach(session => {
        const trainingSessionButton = document.createElement('button');
        trainingSessionButton.className = 'session-button';
        trainingSessionButton.textContent = `${session}`;
        trainingSessionButton.onclick = () => joinGame(session);
        container.appendChild(trainingSessionButton);
    });
}

async function renderResult(response) {
    if (response.error) {
        const errorMessage = response.error.detail || 'An unknown error occurred.';
        const resultsHTML = `
            <h3>Error</h3>
            <p>${errorMessage}</p>
        `;
        resultsDiv.innerHTML = resultsHTML;
        return;
    }

    console.log(response);
    const {turn, tasks, expirationTime} = response;
    const {playerName, score, bonusPoints, dimension, taskViolations} = turn;

    // Fetch rule descriptions
    const ruleDescriptions = await fetchRuleDescriptions();

    // Map violations to descriptions
    const describedTasks = mapTasksToRulesDescriptions(tasks, ruleDescriptions);

    // Construct the results HTML
    let headerHTML = `
        <h3>Player: ${playerName}</h3>
        <p>Score: ${score}</p>
        <p>Bonus Points: ${bonusPoints ? "Yes" : "No"}</p>
        `;

    let violationHTML = '';
    if (taskViolations && taskViolations.length > 0) {
        violationHTML = `
            <h4>Task Violations:</h4>
            <ul>
                ${taskViolations.map(taskViolation => `<li>${taskViolation}</li>`).join('')}
            </ul>
            `;
    }

    let explainationHTML = '';
    if (describedTasks && describedTasks.length > 0) {
        explainationHTML = `
            <h4>Task Explanations:</h4>
            <ul>
                ${describedTasks.map(describedTask => `<li>${describedTask}</li>`).join('')}
            </ul>
            <p>Expiration Time: ${new Date(expirationTime).toLocaleString()}</p>
        `;
    }

    let resultsHTML = headerHTML;
    resultsHTML += violationHTML;
    resultsHTML += explainationHTML;

    // Display the results
    resultsDiv.innerHTML = resultsHTML;
}


function renderStatistics(tasks, turns, expirationTime) {
    const container = document.getElementById('training-section');

    // Clear any existing content
    container.innerHTML = '<h2>Game Statistics</h2>';

    // // Create a section for tasks
    // const tasksSection = document.createElement('div');
    // tasksSection.innerHTML = '<h3>Tasks</h3>';
    // const tasksList = document.createElement('ul');
    // tasks.forEach(task => {
    //     const taskItem = document.createElement('li');
    //     taskItem.textContent = task;
    //     tasksList.appendChild(taskItem);
    // });
    // tasksSection.appendChild(tasksList);
    // container.appendChild(tasksSection);

    // Create a section for turns
    if (turns && Array.isArray(turns)) {
        const turnsSection = document.createElement('div');
        turnsSection.innerHTML = '<h3>Turns</h3>';
        const turnsTable = document.createElement('table');
        turnsTable.innerHTML = `
            <tr>
                <th>Player Name</th>
                <th>Score</th>
                <th>Bonus Points</th>
<!--                <th>Dimension</th>-->
<!--                <th>Task Violations</th>-->
            </tr>
        `;
        turns.forEach(turn => {
            const turnRow = document.createElement('tr');
            turnRow.innerHTML = `
                <td>${turn.playerName}</td>
                <td>${turn.score}</td>
                <td>${turn.bonusPoints}</td>
                `;
            // <td>${Object.entries(turn.dimension).map(([key, value]) => `${key}: ${value}`).join(', ')}</td>
            // <td>${turn.taskViolations ? turn.taskViolations.join(', ') : 'None'}</td>

            turnsTable.appendChild(turnRow);
        });
        turnsSection.appendChild(turnsTable);
        container.appendChild(turnsSection);
    } else {
        console.error('Turns data is not available or not an array.');
    }

    // Create a section for expiration time
    if (expirationTime) {
        const expirationSection = document.createElement('div');
        expirationSection.innerHTML = `<h3>Expiration Time</h3><p>${new Date(expirationTime).toLocaleString()}</p>`;
        container.appendChild(expirationSection);
    } else {
        console.error('Expiration time data is not available.');
    }
}


