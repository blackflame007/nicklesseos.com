window.toggleTheme = function toggleTheme() {
    const themes = ["night", "coffee", "synthwave", "halloween", "sunset"];
    const currentTheme = document.body.getAttribute('data-theme') as string;
    const currentIndex = themes.indexOf(currentTheme);
    const nextIndex = (currentIndex + 1) % themes.length;
    document.body.setAttribute('data-theme', themes[nextIndex]);
}