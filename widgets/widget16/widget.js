document.getElementById('email-button').addEventListener('click', function () {
    console.log('Notify button clicked');
    const email = document.getElementById('email-input').value;
    if (validateEmail(email)) {
        alert('������� �� ��� ������! �� �������� ��� � ����� ������������.');
        ////////////////////////////////////////////
        sessionStorage.setItem('save_widgetState.notActionWidget', "false");
        sessionStorage.removeItem('sendwidget');
        sessionStorage.setItem('sendwidget', true);
        ////////////////////////////////////////////

        document.getElementById('notification-widget').style.display = 'none';
        // ����� ����� �������� ������ ��� �������� email �� ������
    } else {
        alert('����������, ������� ���������� email �����.');
    }
});

function validateEmail(email) {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(String(email).toLowerCase());
}