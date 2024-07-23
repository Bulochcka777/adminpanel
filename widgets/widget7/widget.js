document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById('callbackForm').addEventListener('submit', function (event) {
    event.preventDefault(); // Предотвращаем стандартное поведение формы

    // Получаем значения полей
    const name = document.getElementById('name').value;
    const phone = document.getElementById('phone').value;

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