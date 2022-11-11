import { OnDestroy, Directive, ElementRef, HostListener, Input } from '@angular/core';

@Directive({
  selector: '[appTooltip]'
})
export class TooltipDirective implements OnDestroy {
  @Input("appTooltip") tooltip = '';
  @Input() delay?= 300;
  timer?: NodeJS.Timeout;
  myPopup?: HTMLDivElement;

  constructor(private el: ElementRef) { }

  ngOnDestroy(): void {
    if (this.myPopup) { this.myPopup.remove() }
  }

  @HostListener('mouseenter') onMouseEnter() {
    this.timer = setTimeout(() => {
      let x = this.el.nativeElement.getBoundingClientRect().left; // Get the middle of the element
      let y = this.el.nativeElement.getBoundingClientRect().top + (16 * 3); // Get the bottom of the element, plus a little extra
      this.createTooltipPopup(x, y);
    }, this.delay)
  }

  @HostListener('mouseleave') onMouseLeave() {
    if (this.timer) clearTimeout(this.timer);
    if (this.myPopup) { this.myPopup.remove() }
  }

  private createTooltipPopup(x: number, y: number) {
    let popup = document.createElement('div');
    popup.innerHTML = this.tooltip;
    popup.setAttribute("class", "bg-gray-800 text-white border-gray border opacity-0 transition-opacity Tooltip-Container p-4 absolute rounded-lg duration-300");
    popup.style.top = "112.5%";
    popup.style.left = "50%";
    popup.style.zIndex = "2";
    popup.style.width = "18rem";
    popup.style.transform = `translate(-50%, 0)`;
    let caretDown = document.createElement('div');
    caretDown.setAttribute("class", "Tooltip-Arrow w-0 h-0 border-x-8 border-b-8 border-b-white absolute top-0 left-1/2")
    caretDown.setAttribute("data-popper-arrow", "")
    caretDown.style.transform = `translate(-50%, -100%)`;
    caretDown.style.borderLeft = "1rem solid transparent";
    caretDown.style.borderRight = "1rem solid transparent";
    caretDown.style.borderBottom = "1rem solid #1f2937";
    popup.appendChild(caretDown)

    // document.body.appendChild(popup);
    this.el.nativeElement.appendChild(popup);
    this.myPopup = popup;
    // setTimeout(() => {
    //   if (this.myPopup) this.myPopup.remove();
    // }, 5000); // Remove tooltip after 5 seconds
  }
}
