function showNotification() {
    if (Notification.permission === "granted") {
        var notification = new Notification("Hello, World!");
    } else if (Notification.permission !== "denied") {
        Notification.requestPermission().then(function(permission) {
            if (permission === "granted") {
                var notification = new Notification("Hello, World!");
            } 
        });
    }
}