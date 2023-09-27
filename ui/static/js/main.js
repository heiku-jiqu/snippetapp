var navLinks = document.querySelectorAll("nav a");
navLinks.forEach(node => {
  if (node.getAttribute('href') == window.location.pathname) {
    node.classList.add("live");
  }
});


