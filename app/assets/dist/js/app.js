tailwind.config = {
  theme: {
    extend: {
      colors: {
        // Dark Theme Colors
        darkBlue: '#2b2e4a',
        brightRed: '#e84545',
        deepRed: '#903749',
        darkPurple: '#53354a',
        // Light Theme Colors
        lightBlue: '#1D5B79',
        lightTeal: '#468B97',
        lightRed: '#EF6262',
        lightOrange: '#F3AA60',
      }
    }
  }
}


function toggleTheme() {
  const body = document.body;
  body.classList.toggle('theme-dark');
  body.classList.toggle('theme-light');

  // Update the header colors when toggling theme
  const isDarkTheme = body.classList.contains('theme-dark');
  document.querySelector('header').className = isDarkTheme ? 'bg-darkBlue text-white' : 'bg-lightBlue text-black';
  document.querySelector('header span').className = isDarkTheme ? 'ml-3 text-xl text-brightRed' : 'ml-3 text-xl text-lightRed';
  document.querySelectorAll('nav a').forEach(link => {
    link.className = isDarkTheme ? 'mr-5 hover:text-deepRed' : 'mr-5 hover:text-lightTeal';
  });
  document.querySelector('header button').className = isDarkTheme ? 'py-2 px-4 bg-brightRed text-white rounded' : 'py-2 px-4 bg-lightOrange text-black rounded';
}
