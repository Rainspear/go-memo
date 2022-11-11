import { OnDestroy, Directive, ElementRef, HostListener, Input } from '@angular/core';

@Directive({
  selector: '[appTooltip]'
})
export class TooltipDirective implements OnDestroy {
  @Input() tooltip = '';
  @Input() delay?= 300;
  timer?: NodeJS.Timeout;
  myPopup?: HTMLDivElement;

  constructor(private el: ElementRef) { }

  ngOnDestroy(): void {
    if (this.myPopup) { this.myPopup.remove() }
  }

  @HostListener('mouseenter') onMouseEnter() {
    this.timer = setTimeout(() => {
      let x = this.el.nativeElement.getBoundingClientRect().left + this.el.nativeElement.offsetWidth / 2; // Get the middle of the element
      let y = this.el.nativeElement.getBoundingClientRect().top + this.el.nativeElement.offsetHeight + 6; // Get the bottom of the element, plus a little extra
      this.createTooltipPopup(x, y);
    }, this.delay)
  }

  @HostListener('mouseleave') onMouseLeave() {
    // if (this.timer) clearTimeout(this.timer);
    // if (this.myPopup) { this.myPopup.remove() }
  }

  private createTooltipPopup(x: number, y: number) {
    let popup = document.createElement('div');
    popup.innerHTML = this.tooltip;
    popup.setAttribute("class", `
    Tooltip-Container po p-4 absolute 
    -translate-y-full inline-block z-10 text-sm font-medium text-white bg-gray-900 rounded-lg shadow-sm opacity-0 transition-opacity duration-300 dark:bg-gray-700
    `);
    popup.style.top = (y - 24).toString() + "px";
    popup.style.left = (x).toString() + "px";
    let caretDown = document.createElement('div');
    caretDown.setAttribute("class", "tooltip-arrow")
    caretDown.setAttribute("data-popper-arrow", "")
    popup.appendChild(caretDown)
    document.body.appendChild(popup);
    this.myPopup = popup;
    // setTimeout(() => {
    //   if (this.myPopup) this.myPopup.remove();
    // }, 5000); // Remove tooltip after 5 seconds
  }
}
