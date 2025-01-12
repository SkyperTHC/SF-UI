import { Component, Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Config } from 'src/app/config/config';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {
  title = 'Segfault';
  activeMenu: String = "terminal"
  sidebarVisible: boolean = true
  sidebarFirstLoad: boolean = true

  menuItems: Array<any> = []

  router!: Router
  desktopRequested: boolean = false


  constructor(router: Router) {
    this.router = router
    this.menuItems.push({ ilink: '../assets/icons/term.svg', name: "terminal" })
    if (!Config.DesktopDisabled) {
      this.menuItems.push({ ilink: '../assets/icons/desk.svg', name: "desktop" })
    }
    this.menuItems.push({ ilink: '../assets/icons/ports.svg', name: "ports" })
    this.menuItems.push({ ilink: '../assets/icons/web.svg', name: "web" })
  }

  setActiveMenu(name: string) {
    this.activeMenu = name
    if (this.activeMenu == "desktop" && !this.desktopRequested) {
      this.desktopRequested = true
    }
  }
}
