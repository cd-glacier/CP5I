//
//  ViewController.swift
//  CP5I
//
//  Created by glacier on 2017/09/12.
//  Copyright © 2017年 Bteam. All rights reserved.
//

import UIKit
import Alamofire
import SwiftyJSON

class ViewController: UIViewController, UISearchBarDelegate, UITableViewDelegate, UITableViewDataSource {

	@IBOutlet weak var searchBar: UISearchBar!
	@IBOutlet weak var label: UILabel!
	@IBOutlet weak var tableView: UITableView!

	var tableData: [String] = []

	override func viewDidLoad() {
		super.viewDidLoad()
		// Do any additional setup after loading the view, typically from a nib.

		searchBar.placeholder = "検索"
		searchBar.delegate = self

		tableView.delegate = self
		tableView.dataSource = self

		Alamofire.request("http://noticeweb.net/api/easy/recipe").responseJSON { response in
			let json = JSON(response.result.value)

			json["data"].forEach{(_, data) in
				self.tableData.append(data["name"].string!)
                self.tableView.reloadData()
			}
		}
	}

	func searchBarSearchButtonClicked(_ searchBar: UISearchBar) {
		label.text = searchBar.text
		//キーボードを閉じる
		self.view.endEditing(true)
	}

	func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
		return tableData.count
	}

	func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
		let cell = UITableViewCell(style: UITableViewCellStyle.default, reuseIdentifier: "myCell")
		cell.textLabel?.text = tableData[indexPath.row]
		return cell
	}

	override func didReceiveMemoryWarning() {
		super.didReceiveMemoryWarning()
		// Dispose of any resources that can be recreated.
	}


}

