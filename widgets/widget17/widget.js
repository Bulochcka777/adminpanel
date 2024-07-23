document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById('obr-zvon-button').addEventListener('click', function () {
    // �������� �������� �����
    const name = document.getElementById('name-input').value;
    const phone = document.getElementById('phone-input').value;

    // ��������� ��������� ������ (����� �������� ����� ������� ���������)
    if (name === '' || phone === '') {
        alert('����������, ��������� ��� ����.');
        return;
    }

    alert('��� ����������!');

    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);

    document.getElementById('notification-widget').style.display = 'none';
});