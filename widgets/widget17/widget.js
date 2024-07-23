document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById('obr-zvon-button').addEventListener('click', function () {
    // Получаем значения полей
    const name = document.getElementById('name-input').value;
    const phone = document.getElementById('phone-input').value;

    // Проверяем введенные данные (можно добавить более сложную валидацию)
    if (name === '' || phone === '') {
        alert('Пожалуйста, заполните все поля.');
        return;
    }

    alert('Вам перезвонят!');

    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);

    document.getElementById('notification-widget').style.display = 'none';
});