//
//  ViewController.swift
//  CP5I
//
//  Created by glacier on 2017/09/12.
//  Copyright © 2017年 Bteam. All rights reserved.
//

import UIKit
import Alamofire
import AlamofireImage
import SwiftyJSON

class ViewController: UIViewController, UISearchBarDelegate, UITableViewDelegate, UITableViewDataSource {

	@IBOutlet weak var searchBar: UISearchBar!
	@IBOutlet weak var tableView: UITableView!

	struct Recipe {
		var name: String
		var imageUrl: String
	}

	var recipes: [Recipe] = []
	var findFood: [String] = []
	var findKitchenware: [String] = []

	override func viewDidLoad() {
		super.viewDidLoad()
		// Do any additional setup after loading the view, typically from a nib.

		searchBar.placeholder = "食材で検索"
		searchBar.delegate = self
		tableView.delegate = self
		tableView.dataSource = self

		req(food: [], kitchenwares: [])
	}
	func searchBarSearchButtonClicked(_ searchBar: UISearchBar) {
		let text:String = searchBar.text!
		findFood = text.components(separatedBy: " ")

		recipes = []	
		req(food: findFood, kitchenwares: findKitchenware)
		//キーボードを閉じる
		self.view.endEditing(true)
	}

	func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
		return recipes.count
	}

	func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
		// セルの高さを設定
		return 100
	}

	func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
		let cell = UITableViewCell(style: .subtitle, reuseIdentifier: "myCell")
		cell.textLabel?.text = recipes[indexPath.row].name
		cell.detailTextLabel?.text = "ここが詳細テキストラベルです"
		cell.accessoryType = UITableViewCellAccessoryType.disclosureIndicator
		cell.imageView?.frame = CGRect(x: 0, y: 0, width: 100, height: 100)
		cell.imageView?.af_setImage(withURL: NSURL(string:recipes[indexPath.row].imageUrl)! as URL,  placeholderImage: UIImage(named: "hoiru.png"), imageTransition: .crossDissolve(0.5))
		
		return cell
	}

	@IBAction func pushFlypanButton(_ sender: UIButton) {
		if findKitchenware.contains(where: { $0 == "フライパン" }) {
			findKitchenware.remove(at: findKitchenware.index(of: "フライパン")!)
		} else {
			findKitchenware.append("フライパン")
		}

		recipes = []
		req(food: findFood, kitchenwares: findKitchenware)
	}

	@IBAction func pushPotButton(_ sender: UIButton) {
		tableView.reloadData()
		print(recipes)
	}

	override func didReceiveMemoryWarning() {
		super.didReceiveMemoryWarning()
		// Dispose of any resources that can be recreated.
	}

	func req(food: [String], kitchenwares: [String]){
		let url:String = "http://noticeweb.net/api/easy/recipe"
		Alamofire.request(url, parameters: ["food": food.joined(separator: ","), "kitchenware": kitchenwares.joined(separator: ",")]).responseJSON { response in
			let json = JSON(response.result.value)
			json["data"].forEach{(_, data) in
				self.recipes.append(Recipe(name: data["name"].string!, imageUrl: data["image_url"].string!))
				self.tableView.reloadData()
			}
		}
	}

}

