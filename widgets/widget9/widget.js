document.getElementById('vk-button').addEventListener('click', function () {
    alert('������� � ��');
    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////
    document.getElementById('notification-widget').style.display = 'none';
    // ����� ����� �������� ������ ��� �������� email �� ������
});

document.getElementById('telegram-button').addEventListener('click', function () {
    alert('������� � ��������');
    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////
    document.getElementById('notification-widget').style.display = 'none';
    // ����� ����� �������� ������ ��� �������� email �� ������
});

document.getElementById('instagram-button').addEventListener('click', function () {
    alert('������� � ���������');
    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////
    document.getElementById('notification-widget').style.display = 'none';
    // ����� ����� �������� ������ ��� �������� email �� ������
});