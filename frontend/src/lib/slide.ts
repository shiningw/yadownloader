class Slide {
    button: HTMLButtonElement;
    container: HTMLElement;
    height: number;
    width: number;
    constructor(btn?: string, contentContainer?: string) {
        this.button = document.querySelector(btn) as HTMLButtonElement;
        if (!this.button) {
            this.button = document.querySelector('[data-apps-slide-toggle]') as HTMLButtonElement;
        }
        if (!contentContainer) {
            contentContainer = this.button.getAttribute('data-apps-slide-toggle');
        }
        this.container = document.querySelector(contentContainer) as HTMLDivElement;
        this.button.addEventListener('click', (e) => {
            if (this.container.classList.contains('apps-slide-open')) {
                this.slideDown();
                this.container.classList.toggle('apps-slide-open');
            } else {
                this.slideUp();
                this.container.classList.toggle('apps-slide-open');
            }
        })
        this.height = this._getContainerHeight();
    }
    static create(btn?: string, contentContainer?: string): Slide {
        return new Slide(btn, contentContainer);
    }
    slideUp(): void {
        this.container.style.transition = "height,padding 400ms ease-in-out";
        this.container.style.height = "0px";
        this.container.style.display = "block";
        this.container.style.boxSizing = "border-box";

        window.setTimeout(() => {
            //this.container.style.removeProperty('height');
            this.container.style.height = this.height + "px";
            this.container.style.removeProperty('overflow');
            this.container.style.removeProperty('transition');
        }, 10);
    }

    slideDown(): void {
        this.container.style.overflow = "hidden";
        this.container.style.transition = "height,padding 400ms ease-in-out";

        window.setTimeout(() => {
            this.container.style.height = "0px";
            this.container.style.padding = "0px";
            this.container.style.removeProperty('overflow');
            this.container.style.removeProperty('transition');
            this.container.style.removeProperty('display');
        }, 10)
    }
    _getContainerHeight(): number {
        const display = window.getComputedStyle(this.container).display;
        this.container.style.display = "block";
        let height = this.container.clientHeight;
        this.container.style.display = "none";
        return height
    }
}

export default Slide