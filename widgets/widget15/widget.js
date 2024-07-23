document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

// ������� ��� ��������� ������ �����-������
function handleRadioSelection() {
    let selectedRadio = document.querySelector('input[name="radio1"]:checked');
    if (!selectedRadio) {
        alert("����������, �������� ���� �� ���������.");
        return;
    }

    let selectedValue = selectedRadio.value;
    console.log('������� �����:', selectedValue);

    // �������� ������
    let data = {
        question1: selectedValue,
    };

    // ������� ��������� ������ � ������� (��� ���������� �� �� ������)
    console.log(data);

    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////

    // �������� ������
    document.getElementById('notification-widget').style.display = 'none';
}

// ��������� ���������� ������� ��� ���� �����-������
document.querySelectorAll('.radio-input').forEach(function (radio) {
    radio.addEventListener('change', handleRadioSelection);
});