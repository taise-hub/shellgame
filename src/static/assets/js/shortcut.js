document.addEventListener('keydown', (event) => {
    var keyName = event.key;
    if (event.ctrlKey) {
      if (keyName == "l") {
          clearLeft()
      }
      else if (keyName == "p"){
          previousCommand()
      }
      else if (keyName == "n"){
          nextCommand()
      }
    }
  });