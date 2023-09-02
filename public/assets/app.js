let theme = 0;
let themes = ["light", "dark"];

document.querySelector("html").setAttribute("data-theme", themes[theme]);

function toggleTheme() {
  theme = (theme + 1) % themes.length;
  document.querySelector("html").setAttribute("data-theme", themes[theme]);
}
