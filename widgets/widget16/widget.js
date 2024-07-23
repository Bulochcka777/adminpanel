document.getElementById('email-button').addEventListener('click', function () {
    console.log('Notify button clicked');
    const email = document.getElementById('email-input').value;
    if (validateEmail(email)) {
        alert('Спасибо за ваш запрос! Мы уведомим вас о новых поступлениях.');
        ////////////////////////////////////////////
        sessionStorage.setItem('save_widgetState.notActionWidget', "false");
        sessionStorage.removeItem('sendwidget');
        sessionStorage.setItem('sendwidget', true);
        ////////////////////////////////////////////

        document.getElementById('notification-widget').style.display = 'none';
        // Здесь можно добавить логику для отправки email на сервер
    } else {
        alert('Пожалуйста, введите корректный email адрес.');
    }
});

function validateEmail(email) {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(String(email).toLowerCase());
}