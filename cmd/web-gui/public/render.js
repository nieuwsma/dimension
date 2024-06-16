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

        const { Name, Description } = task;
        const parts = Name.split('-');

        const card = document.createElement('div');
        card.className = 'card';
        const taskName = document.createElement('div');
        taskName.className = 'task-name';
        taskName.textContent = Name;
        card.appendChild(taskName);

        // Create a graphical representation based on the task type
        switch(parts[0]) {
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
    switch(colorCode) {
        case 'G': return 'green';
        case 'B': return 'blue';
        case 'K': return 'black';
        case 'O': return 'orange';
        case 'W': return 'white';
        default: return 'gray';
    }
}

function renderSessions(sessions) {
    const container = document.getElementById('session-list');
    container.innerHTML = ''; // Clear previous content

    sessions.forEach(session => {
        const sessionDiv = document.createElement('div');
        sessionDiv.className = 'session';
        sessionDiv.innerHTML = `Session ID: ${session} <button onclick="joinSession('${session}')">Join</button>`;
        container.appendChild(sessionDiv);
    });
}
