document.getElementById('skidka-button').addEventListener('click', function () {
    alert('��� ������ ���� ������� �� �������� � ������');
    ////////////////////////////////////////////
    sessionStorage.setItem('save_widgetState.notActionWidget', "false");
    sessionStorage.removeItem('sendwidget');
    sessionStorage.setItem('sendwidget', true);
    ////////////////////////////////////////////
    document.getElementById('notification-widget').style.display = 'none';
    // ����� ����� �������� ������ ��� �������� email �� ������
});