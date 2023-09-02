let theme = 0;
let themes = ["dark", "light"];

document.querySelector("html").setAttribute("data-theme", themes[theme]);

function toggleTheme() {
  theme = (theme + 1) % themes.length;
  document.querySelector("html").setAttribute("data-theme", themes[theme]);
}
