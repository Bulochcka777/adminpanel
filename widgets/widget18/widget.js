document.getElementById('skidka-button').addEventListener('click', function () {
    alert('Тут должен быть переход на страницу с акцией');
    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////
    document.getElementById('notification-widget').style.display = 'none';
    // Здесь можно добавить логику для отправки email на сервер
});