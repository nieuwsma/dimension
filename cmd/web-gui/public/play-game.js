document.addEventListener('DOMContentLoaded', async () => {
    const urlParams = new URLSearchParams(window.location.search);
    const trainID = urlParams.get('trainID');
    console.log('trainID:', trainID); // Debugging line to log trainID

    if (trainID) {
        document.getElementById('train-id').textContent = `Training Session ID: ${trainID}`;
        const response = await fetchTrainingSession(trainID);  // Fetch tasks for the specific trainID
        if (response && response.tasks) {
            const tasks = response.tasks;
            console.log('Tasks:', tasks);  // Debugging line to log the tasks
            renderTasks(tasks);
        } else {
            console.error('Failed to fetch tasks for the training session.');
        }
    } else {
        console.error('No  trainID found.');
    }
});

const slots = document.querySelectorAll('.slot');
const colorPicker = document.getElementById('color-picker');
const spheres = document.querySelectorAll('.available-sphere');
const playerNameInput = document.getElementById('player-name');
const resultsDiv = document.getElementById('result-content');

let currentSlot = null;

document.querySelectorAll('.sphere').forEach(slot => {
    slot.addEventListener('click', (event) => {
        currentSlot = event.target;
        const rect = currentSlot.getBoundingClientRect();
        colorPicker.style.top = `${rect.bottom + window.scrollY}px`;
        colorPicker.style.left = `${rect.left + window.scrollX}px`;
        colorPicker.style.display = 'block';
    });
});

document.querySelectorAll('.color-option, .trash-option').forEach(option => {
    option.addEventListener('click', () => {
        const color = option.getAttribute('data-color');

        if (option.classList.contains('trash-option')) {
            resetSlot(currentSlot);
        } else if (currentSlot.getAttribute('data-color') === color) {
            resetSlot(currentSlot);
        } else if (!option.classList.contains('line-through')) {
            setColor(currentSlot, color, option.classList[1]);
        }

        colorPicker.style.display = 'none';
        updateColorPicker();
    });
});

function setColor(slot, color, colorName) {
    const previousColor = slot.getAttribute('data-color');

    if (previousColor) {
        const availableSphere = document.querySelector(`.available-sphere.${previousColor}-sphere.line-through`);
        if (availableSphere) {
            availableSphere.classList.remove('line-through');
        }
    }

    slot.style.background = colorName;
    slot.style.backgroundColor = colorName;
    slot.style.backgroundImage = "none";
    slot.setAttribute('data-color', color);

    if (color === 'g' || color === 'k' || color === 'b') {
        slot.style.color = 'white';
    } else {
        slot.style.color = 'black';
    }

    updateColorPicker();
}

function resetSlot(slot) {
    const previousColor = slot.getAttribute('data-color');

    if (previousColor) {
        const availableSphere = document.querySelector(`.available-sphere.${previousColor}-sphere.line-through`);
        if (availableSphere) {
            availableSphere.classList.remove('line-through');
        }
    }

    slot.style.background = "repeating-linear-gradient(45deg, #ffffff, #ffffff 10px, #f0f0f0 10px, #f0f0f0 20px)";
    slot.style.color = 'black';
    slot.removeAttribute('data-color');

    updateColorPicker();
}

function updateColorPicker() {
    const colorCounts = {
        'g': 0,
        'k': 0,
        'b': 0,
        'w': 0,
        'o': 0
    };

    document.querySelectorAll('.sphere[data-color]').forEach(slot => {
        const color = slot.getAttribute('data-color');
        if (color) {
            colorCounts[color]++;
        }
    });

    document.querySelectorAll('.color-option').forEach(option => {
        const color = option.getAttribute('data-color');
        if (colorCounts[color] >= 3) {
            option.classList.add('line-through');
        } else {
            option.classList.remove('line-through');
        }
    });

    document.querySelectorAll('.available-sphere').forEach(sphere => {
        const color = sphere.getAttribute('data-color');
        const sphereCount = colorCounts[color];
        if (sphereCount > 0) {
            sphere.classList.add('line-through');
            colorCounts[color]--;
        } else {
            sphere.classList.remove('line-through');
        }
    });
}

document.getElementById('submit-btn').addEventListener('click', async () => {
    console.log('Submit button clicked');  // Debugging line to check button click
    const playerName = playerNameInput.value.trim();
    if (!playerName) {
        alert('Player name is required.');
        return;
    }

    const urlParams = new URLSearchParams(window.location.search);  // Ensure urlParams is declared here
    const trainID = urlParams.get('trainID');
    console.log('Train ID:', trainID);  // Debugging line to log trainID
    console.log('Player Name:', playerName);  // Debugging line to log playerName

    const slotData = {};
    document.querySelectorAll('.sphere').forEach(slot => {
        const slotId = slot.getAttribute('id');
        const color = slot.getAttribute('data-color');
        if (color) {
            slotData[slotId] = color;
        }
    });

    if (validateSubmission(slotData)) {
        const response = await submitTurn(trainID, playerName, slotData);
        await renderResult(response);
    } else {
        alert('Validation failed: A color can be used only 3 times.');
    }
});

document.getElementById('reset-btn').addEventListener('click', () => {
    resetAll();
});

function resetAll() {
    document.querySelectorAll('.sphere').forEach(slot => {
        resetSlot(slot);
    });
}

function validateSubmission(slotData) {
    const colorCounts = {};
    Object.values(slotData).forEach(color => {
        colorCounts[color] = (colorCounts[color] || 0) + 1;
    });
    return !Object.values(colorCounts).some(count => count > 3);
}

function mapTasksToRulesDescriptions(tasks, ruleDescriptions) {
    return tasks.map(task => {
        const rule = ruleDescriptions.find(r => task.includes(r.Name));
        return rule ? rule.Description : task;
    });
}


// Hide color picker when clicking outside
document.addEventListener('click', (event) => {
    if (!colorPicker.contains(event.target) && !event.target.classList.contains('sphere')) {
        colorPicker.style.display = 'none';
    }
});

updateColorPicker();
