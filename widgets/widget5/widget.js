document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById('notify-button').addEventListener('click', function () {
    const phone = document.getElementById('phone-input').value;
    if (validatePhone(phone)) {
        alert('Спасибо за ваш запрос! Мы уведомим вас.');
        document.getElementById('notification-widget').style.display = 'none';
        // Здесь можно добавить логику для отправки номера телефона на сервер
    } else {
        alert('Пожалуйста, введите корректный номер телефона.');
    }
});

function validatePhone(phone) {
    const re = /^\d{1,4} \d{3} \d{3} \d{2} \d{2}$/;
    return re.test(String(phone).toLowerCase());
}