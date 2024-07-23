document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById("reaction-button").addEventListener("click", function () {
    // Получаем выбранные значения radio-кнопок для каждого вопроса
    let question1Value = document.querySelector('input[name="reaktion1"]:checked');
    let question2Value = document.querySelector('input[name="reaktion2"]:checked');

    // Проверяем, выбраны ли все radio-кнопки
    if (!question1Value || !question2Value) {
        alert("Пожалуйста, ответьте на все вопросы.");
        return;
    }

    // Собираем значения
    let data = {
        question1: question1Value.value,
        question2: question2Value.value
    };

    // Выводим собранные данные в консоль (или отправляем их на сервер)
    console.log(data);

    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////

    document.getElementById('notification-widget').style.display = 'none';
});