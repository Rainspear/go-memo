import { HostBinding, ElementRef, Directive, HostListener } from '@angular/core';

@Directive({
  selector: '[appDropdown]'
})
export class DropdownDirective {
  @HostBinding('class.open') classes = false;

  constructor(private elementRef: ElementRef) { }
  
  @HostListener('document:click', ['$event']) toggleOpen(event: Event) {
    this.classes = this.elementRef.nativeElement.contains(event.target) ? !this.classes : false;
  }
  // @HostListener('click') onClick(eventData: Event) {
  //   this.classes = !this.classes;
  // }
}
