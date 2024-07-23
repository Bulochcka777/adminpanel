document.getElementById('close-widget1').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById('close-widget2').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById("widget1").addEventListener("click", function () {
    const widget1 = document.getElementById('widget1');
    const widget2 = document.getElementById('widget2');

    widget1.addEventListener('click', function () {
        widget1.classList.add('hidden');
        setTimeout(function () {
            widget2.classList.remove('hidden');
        }, 500); // Задержка совпадает с продолжительностью анимации
    });
});