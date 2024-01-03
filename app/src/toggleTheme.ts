window.toggleTheme = function toggleTheme() {
    const themes = ["night", "coffee", "synthwave", "halloween", "sunset"];
    const currentTheme = document.body.getAttribute('data-theme') || localStorage.getItem('theme') || 'night';
    const currentIndex = themes.indexOf(currentTheme);
    const nextIndex = (currentIndex + 1) % themes.length;
    document.body.setAttribute('data-theme', themes[nextIndex]);
    localStorage.setItem('theme', themes[nextIndex]);
}

// Apply the theme when the page loads
document.addEventListener('DOMContentLoaded', () => {
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme) {
        document.body.setAttribute('data-theme', savedTheme);
    }
});

