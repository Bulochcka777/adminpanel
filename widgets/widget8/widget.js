document.getElementById('close-widget').addEventListener('click', function () {
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('clousewidget');
    sessionStorage.setItem('clousewidget', true);
    document.getElementById('notification-widget').style.display = 'none';
});

document.getElementById('rasprodasha-button').addEventListener('click', function () {
    alert('��� ������ ���� ������� �� �������� � ������');
    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////
    document.getElementById('notification-widget').style.display = 'none';
    // ����� ����� �������� ������ ��� �������� email �� ������
});