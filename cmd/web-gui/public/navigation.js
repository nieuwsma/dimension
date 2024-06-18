function homePage() {
    window.location.href = 'index.html';
}

function gameStatistics() {
    const urlParams = new URLSearchParams(window.location.search);
    const trainID = urlParams.get('trainID');
    window.location.href = `game-statistics.html?trainID=${trainID}`;
}

function playGame() {
    const urlParams = new URLSearchParams(window.location.search);
    const trainID = urlParams.get('trainID');
    window.location.href = `play-game.html?trainID=${trainID}`;
}

function viewRules() {
    window.location.href = 'rules.html';
}