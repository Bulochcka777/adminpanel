document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById("test-button").addEventListener("click", function () {
    // �������� ��������� �������� radio-������ ��� ������� �������
    let question1Value = document.querySelector('input[name="radio1"]:checked');
    let question2Value = document.querySelector('input[name="radio2"]:checked');
    let question3Value = document.querySelector('input[name="radio3"]:checked');

    // ���������, ������� �� ��� radio-������
    if (!question1Value || !question2Value || !question3Value) {
        alert("����������, �������� �� ��� �������.");
        return;
    }

    // �������� ��������
    let data = {
        question1: question1Value.value,
        question2: question2Value.value,
        question3: question3Value.value
    };

    // ������� ��������� ������ � ������� (��� ���������� �� �� ������)
    console.log(data);

    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////

    document.getElementById('notification-widget').style.display = 'none';
});