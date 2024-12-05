// Presentation/Assets/js/htmx-extensions.js

htmx.defineExtension('errorResponseSwap', {    
    onEvent: function(name, evt) {
        if (name === "htmx:beforeSwap") {
            console.log("coucou")
            var xhr = evt.detail.xhr;
            if (xhr.status >= 400 && xhr.status < 600) {
                // Forcer HTMX à swap le contenu même en cas d'erreur
                evt.detail.shouldSwap = true;
                // Optionnel : définir le style de swap si nécessaire
                // evt.detail.swapStyle = 'outerHTML';
            }
        }
    }
});