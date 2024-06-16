document.addEventListener('DOMContentLoaded', async () => {
    const rules = await fetchRules();
    console.log('Rules:', rules);  // Debugging line to log the rules
    renderTasks(rules, 'rules-container', true);
});

function goBack() {
    window.location.href = 'index.html';
}
