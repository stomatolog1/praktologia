const API = {
    async get(path) {
        const response = await fetch(`/api/${path}`);
        if (!response.ok) {
            const error = await response.json().catch(() => ({}));
            throw new Error(error.error || `Ошибка запроса: ${response.status}`);
        }
        return response.json();
    },

    async post(path, body) {
        const response = await fetch(`/api/${path}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(body)
        });
        if (!response.ok) {
            const error = await response.json().catch(() => ({}));
            throw new Error(error.error || `Ошибка запроса: ${response.status}`);
        }
        return response.json();
    }
};

const Auth = {
    storageKey: 'praktologia_user',

    getUser() {
        try {
            return JSON.parse(localStorage.getItem(this.storageKey) || 'null');
        } catch {
            return null;
        }
    },

    setUser(user) {
        localStorage.setItem(this.storageKey, JSON.stringify(user));
    },

    logout() {
        localStorage.removeItem(this.storageKey);
    },

    requireAuth(redirectTo = '/static/login.html') {
        const user = this.getUser();
        if (!user) {
            window.location.href = redirectTo;
            return null;
        }
        return user;
    }
};

function getQueryParam(name) {
    return new URLSearchParams(window.location.search).get(name);
}

function showAlert(elementId, message, type) {
    const alert = document.getElementById(elementId);
    if (!alert) return;
    alert.textContent = message;
    alert.className = 'alert ' + type;
    setTimeout(() => {
        alert.className = 'alert';
    }, 4000);
}

function renderEmptyList(containerId, message) {
    const container = document.getElementById(containerId);
    if (!container) return;
    container.innerHTML = `<div class="empty-state"><p>${message}</p></div>`;
}

function formatFio(item) {
    return [item.Surname, item.Name, item.SecondName].filter(Boolean).join(' ');
}

function formatMoney(value) {
    return new Intl.NumberFormat('ru-RU', {
        style: 'currency',
        currency: 'RUB',
        maximumFractionDigits: 2
    }).format(value || 0);
}

function calcSotrudnikCost(sotrudnik, days) {
    const salary = sotrudnik.PayMount || 0;
    const taxRate = (sotrudnik.Tax || 0) / 100;
    const monthlyCost = salary + salary * taxRate;
    const workingDaysPerMonth = 22;
    const dailyCost = monthlyCost / workingDaysPerMonth;
    return days * dailyCost;
}

function calcExecutorCost(executor, quantity) {
    const base = quantity * (executor.Cost || 0);
    const taxRate = executor.TypeDesign === 'НПД' ? 0 : (executor.Tax || 0) / 100;
    return base + base * taxRate;
}

function calcEquipmentCost(equipment, quantity) {
    return quantity * (equipment.Cost || 0);
}

function renderUserBar(containerId) {
    const container = document.getElementById(containerId);
    if (!container) return;

    const user = Auth.getUser();
    if (!user) {
        container.innerHTML = `
            <a href="/static/login.html" class="nav-link">Войти</a>
            <a href="/static/user-create.html" class="nav-link">Создать пользователя</a>
        `;
        return;
    }

    container.innerHTML = `
        <span class="user-badge">Пользователь: ${user.Login}</span>
        <button type="button" class="btn btn-secondary btn-sm" id="logout-btn">Выйти</button>
    `;

    document.getElementById('logout-btn').addEventListener('click', () => {
        Auth.logout();
        window.location.href = '/';
    });
}
