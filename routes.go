package main

import (
	"net/http"

	"github.com/jritsema/gotoolbox/web"
)

// Delete -> DELETE /client/{id} -> delete, clients.html

// Edit   -> GET /client/edit/{id} -> cliente/row-edit.html
// Save   ->   PUT /client/{id} -> update, cliente/row.html
// Cancel ->	 GET /client/{id} -> nothing, cliente/row.html

// Add    -> GET /client/add/ -> clients-add.html (target body with cliente/row-add.html and cliente/row.html)
// Save   ->   POST /client -> add, clients.html (target body without cliente/row-add.html)
// Cancel ->	 GET /client -> nothing, clients.html

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "index.html", dataClient, nil)
}

// GET /client/add
func clientAdd(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "cliente/client-add.html", dataClient, nil)
}

// /GET client/edit/{id}
func clientEdit(r *http.Request) *web.Response {
	id, _ := web.PathLast(r)
	row := getClientByID(id)
	return web.HTML(http.StatusOK, html, "cliente/row-edit.html", row, nil)
}

func projectAdd(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "project/project-add.html", dataProject, nil)
}

func projectEdit(r *http.Request) *web.Response {
	id, _ := web.PathLast(r)
	row := getProjectByID(id)
	return web.HTML(http.StatusOK, html, "project/row-edit.html", row, nil)
}

// GET /client
// GET /client/{id}
// DELETE /client/{id}
// PUT /client/{id}
// POST /client
func clients(r *http.Request) *web.Response {
	id, segments := web.PathLast(r)
	switch r.Method {

	case http.MethodDelete:
		deleteClient(id)
		return web.HTML(http.StatusOK, html, "cliente/cliente.html", dataClient, nil)

	//cancel
	case http.MethodGet:
		if segments > 1 {
			//cancel edit
			row := getClientByID(id)
			return web.HTML(http.StatusOK, html, "cliente/row.html", row, nil)
		} else {
			//cancel add
			return web.HTML(http.StatusOK, html, "cliente/cliente.html", dataClient, nil)
		}

	//save edit
	case http.MethodPut:
		row := getClientByID(id)
		r.ParseForm()
		row.Client = r.Form.Get("client")
		row.Alias = r.Form.Get("alias")
		updateClient(row)
		return web.HTML(http.StatusOK, html, "cliente/row.html", row, nil)

	//save add
	case http.MethodPost:
		row := Client{}
		r.ParseForm()
		row.Client = r.Form.Get("client")
		row.Alias = r.Form.Get("alias")
		addClient(row)
		return web.HTML(http.StatusOK, html, "cliente/cliente.html", dataClient, nil)
	}

	return web.Empty(http.StatusNotImplemented)
}

func projects(r *http.Request) *web.Response {
	id, segments := web.PathLast(r)
	switch r.Method {

	case http.MethodDelete:
		deleteProject(id)
		return web.HTML(http.StatusOK, html, "project/project.html", dataProject, nil)

	//cancel
	case http.MethodGet:
		if segments > 1 {
			//cancel edit
			row := getProjectByID(id)
			return web.HTML(http.StatusOK, html, "project/row.html", row, nil)
		} else {
			//cancel add
			return web.HTML(http.StatusOK, html, "project/project.html", dataProject, nil)
		}

	//save edit
	case http.MethodPut:
		row := getProjectByID(id)
		r.ParseForm()
		row.Project = r.Form.Get("project")
		row.Alias = r.Form.Get("alias")
		row.Client_Id = r.Form.Get("client_id")
		updateProject(row)
		return web.HTML(http.StatusOK, html, "project/row.html", row, nil)

	//save add
	case http.MethodPost:
		row := Project{}
		r.ParseForm()
		row.Project = r.Form.Get("project")
		row.Alias = r.Form.Get("alias")
		row.Client_Id = r.Form.Get("client_id")
		addProject(row)
		return web.HTML(http.StatusOK, html, "project/project.html", dataProject, nil)
	}

	return web.Empty(http.StatusNotImplemented)
}
