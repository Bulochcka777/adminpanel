document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

// Функция для обработки выбора радио-кнопок
function handleRadioSelection() {
    let selectedRadio = document.querySelector('input[name="radio1"]:checked');
    if (!selectedRadio) {
        alert("Пожалуйста, выберите один из вариантов.");
        return;
    }

    let selectedValue = selectedRadio.value;
    console.log('Выбрана опция:', selectedValue);

    // Собираем данные
    let data = {
        question1: selectedValue,
    };

    // Выводим собранные данные в консоль (или отправляем их на сервер)
    console.log(data);

    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////

    // Скрываем виджет
    document.getElementById('notification-widget').style.display = 'none';
}

// Добавляем обработчик события для всех радио-кнопок
document.querySelectorAll('.radio-input').forEach(function (radio) {
    radio.addEventListener('change', handleRadioSelection);
});