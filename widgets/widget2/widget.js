document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById('callbackForm').addEventListener('submit', function (event) {
    event.preventDefault(); // ������������� ����������� ��������� �����

    // �������� �������� �����
    const name = document.getElementById('name').value;
    const phone = document.getElementById('phone').value;

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