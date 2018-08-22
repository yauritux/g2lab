package com.yauritux.model.entity;

import java.util.ArrayList;
import java.util.List;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.Table;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Entity
@Table(name = "zones")
public class Zone {

	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
	private Long id;
	
	@Column(name = "zone_name", nullable = false, unique = true)
	private String zoneName;

	/*
	@ManyToMany(mappedBy = "zones")
	private List<Station> stations = new ArrayList<>();
	*/
	
	@ManyToMany(cascade = { CascadeType.PERSIST, CascadeType.MERGE })
	@JoinTable(name = "zone_fares", joinColumns = @JoinColumn(name = "zone_id"), 
			inverseJoinColumns = @JoinColumn(name = "fare_id"))
	private List<Fare> fares = new ArrayList<>();	
	
	public Long getId() {
		return id;
	}
	
	protected void setId(Long id) {
		this.id = id;
	}
	
	public String getZoneName() {
		return zoneName;
	}
	
	public void setZoneName(String zoneName) {
		this.zoneName = zoneName;
	}
	
	/*
	public List<Station> getStations() {
		return stations;
	}
	
	public void setStations(List<Station> stations) {
		this.stations = stations;
	}
	*/
	
	public List<Fare> getFares() {
		return fares;
	}
	
	public void setFares(List<Fare> fares) {
		this.fares = fares;
	}
	
	public void addFare(Fare fare) {
		fares.add(fare);
		//fare.getZones().add(this);
	}
	
	public void removeFare(Fare fare) {
		fares.remove(fare);
		//fare.getZones().remove(this);
	}	

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((id == null) ? 0 : id.hashCode());
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		Zone other = (Zone) obj;
		if (id == null) {
			if (other.id != null)
				return false;
		} else if (!id.equals(other.id))
			return false;
		return true;
	}
}
