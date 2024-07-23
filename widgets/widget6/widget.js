document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById("reaction-button").addEventListener("click", function () {
    // �������� ��������� �������� radio-������ ��� ������� �������
    let question1Value = document.querySelector('input[name="reaktion1"]:checked');
    let question2Value = document.querySelector('input[name="reaktion2"]:checked');

    // ���������, ������� �� ��� radio-������
    if (!question1Value || !question2Value) {
        alert("����������, �������� �� ��� �������.");
        return;
    }

    // �������� ��������
    let data = {
        question1: question1Value.value,
        question2: question2Value.value
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