
function toggleTheme() {
    const body = document.body;
    const currentTheme = body.getAttribute('data-theme');
    const newTheme = currentTheme === 'light' ? 'dark' : 'light';
    body.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
    updateThemeButton(newTheme);
}

function updateThemeButton(theme) {
    const btn = document.querySelector('.theme-toggle');
    const prismTheme = document.getElementById('prism-theme');
    
    if (btn) {
        btn.innerHTML = theme === 'light' ? '🌙 Dark' : '☀️ Light';
    }
    
    if (prismTheme) {
        // Light mode uses a lighter Prism theme, Dark mode uses Tomorrow Night
        prismTheme.href = theme === 'light' 
            ? 'https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/themes/prism.min.css'
            : 'https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/themes/prism-tomorrow.min.css';
    }
}

// Language Picker Logic
function setupLocalePicker() {
    const picker = document.getElementById('localePicker');
    if (!picker) return;

    const currentPath = window.location.pathname;
    let currentLocale = 'en';
    if (currentPath.includes('index_pt_br.html')) currentLocale = 'pt-BR';
    if (currentPath.includes('index_es.html')) currentLocale = 'es';

    picker.value = currentLocale;
    picker.addEventListener('change', (e) => {
        const target = e.target.value;
        let newPath = 'index.html';
        
        if (target === 'pt-BR') newPath = 'index_pt_br.html';
        else if (target === 'es') newPath = 'index_es.html';
        
        window.location.href = newPath;
    });
}

// Scroll Spy for Navigation
function setupScrollSpy() {
    const sections = document.querySelectorAll('section[id]');
    const navLinks = document.querySelectorAll('.category-pill');

    window.addEventListener('scroll', () => {
        let current = '';
        sections.forEach(section => {
            const sectionTop = section.offsetTop;
            if (pageYOffset >= sectionTop - 150) {
                current = section.getAttribute('id');
            }
        });

        navLinks.forEach(link => {
            link.classList.remove('active');
            if (link.getAttribute('href') === `#${current}`) {
                link.classList.add('active');
            }
        });
    });
}

// Lightbox Logic
function setupLightbox() {
    // Create modal element if not exists
    let modal = document.getElementById('imageModal');
    if (!modal) {
        modal = document.createElement('div');
        modal.id = 'imageModal';
        modal.className = 'modal';
        modal.innerHTML = `
            <span class="close">&times;</span>
            <img class="modal-content" id="modalImg">
            <div id="caption"></div>
        `;
        document.body.appendChild(modal);
    }

    const modalImg = document.getElementById("modalImg");
    const captionText = document.getElementById("caption");
    const span = document.querySelector(".close");

    // Add click event to all diagrams (images in cards or sections)
    document.querySelectorAll('.card img, section > img').forEach(img => {
        // Skip small icons like logo if necessary, but here we want all diagrams
        if (img.closest('.logo')) return;

        img.onclick = function() {
            modal.style.display = "block";
            modalImg.src = this.src;
            captionText.innerHTML = this.alt;
            document.body.style.overflow = 'hidden'; // Disable scroll
        }
    });

    // Close logic
    const closeModal = () => {
        modal.style.display = "none";
        document.body.style.overflow = 'auto'; // Re-enable scroll
    };

    span.onclick = closeModal;
    modal.onclick = (e) => { if (e.target === modal) closeModal(); };
    document.addEventListener('keydown', (e) => { if (e.key === 'Escape') closeModal(); });
}

document.addEventListener('DOMContentLoaded', () => {
    // Load saved theme
    const savedTheme = localStorage.getItem('theme') || 'dark';
    document.body.setAttribute('data-theme', savedTheme);
    updateThemeButton(savedTheme);

    setupLocalePicker();
    setupScrollSpy();
    setupLightbox();

    // Smooth scroll
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            const target = document.querySelector(this.getAttribute('href'));
            if (target) {
                window.scrollTo({
                    top: target.offsetTop - 120,
                    behavior: 'smooth'
                });
            }
        });
    });
});
