document.getElementById('vk-button').addEventListener('click', function () {
    alert('Переход в ВК');
    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////
    document.getElementById('notification-widget').style.display = 'none';
    // Здесь можно добавить логику для отправки email на сервер
});

document.getElementById('telegram-button').addEventListener('click', function () {
    alert('Переход в телеграм');
    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////
    document.getElementById('notification-widget').style.display = 'none';
    // Здесь можно добавить логику для отправки email на сервер
});

document.getElementById('instagram-button').addEventListener('click', function () {
    alert('Переход в инстаграм');
    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////
    document.getElementById('notification-widget').style.display = 'none';
    // Здесь можно добавить логику для отправки email на сервер
});