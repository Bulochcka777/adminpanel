document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById('image-section').addEventListener('click', function (event) {
    const target = event.target;
    if (target.tagName.toLowerCase() === 'img') {
        const link = target.parentElement;
        if (link.tagName.toLowerCase() === 'a') {
            event.preventDefault();
            console.log(`Clicked on image: ${target.alt}`);
            sessionStorage.setItem('save_widgetState.notActionWidget', "false");
            sessionStorage.removeItem('sendwidget');
            sessionStorage.setItem('sendwidget', true);

            document.getElementById('notification-widget').style.display = 'none';

            window.location.href = link.href;
        }
    }
});