function ShowNotification() {
    if (Notification.permission === "granted") {
        new Notification("Hello, World!");
    } else if (Notification.permission !== "denied") {
        Notification.requestPermission().then(function(permission) {
            if (permission === "granted") {
                new Notification("Hello, World!");
            } 
        });
    }
}

function createTodo() {
    if (event.keyCode === 13 || event.key === 'Enter') {
    fetch('/app', {
      method: 'POST',
    })
    new Notification("Todo Created")
      .catch(error => {
        console.error('Error:', error);
      });
    }
    else 
        throw error
}
  
