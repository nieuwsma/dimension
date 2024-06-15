function renderTasks(tasks) {
    if (!tasks || tasks.length === 0) {
        console.error('No tasks to render.');
        return;
    }

    const container = document.getElementById('card-container');
    container.innerHTML = ''; // Clear previous content

    tasks.forEach(task => {
        const card = document.createElement('div');
        card.className = 'card';
        const taskName = document.createElement('div');
        taskName.className = 'task-name';
        taskName.textContent = task;
        card.appendChild(taskName);
        const parts = task.split('-');

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
                card.innerHTML = `<p>Unknown Task: ${task}</p>`;
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
