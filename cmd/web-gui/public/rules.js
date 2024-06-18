document.addEventListener('DOMContentLoaded', async () => {
    const rules = await fetchRules();
    console.log('Rules:', rules);  // Debugging line to log the rules
    renderTasks(rules, 'rules-container', true);
});
